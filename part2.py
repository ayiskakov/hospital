from sqlalchemy import create_engine, text
from sqlalchemy.orm import sessionmaker

# Create a connection to the database
engine = None
try: 
    engine = create_engine(
        'postgresql://postgres:secret@localhost:15432/hospital?sslmode=disable',
    )
    engine.connect()
except Exception as e:
    print(e)
    exit(1)


session = sessionmaker(autocommit=False, autoflush=False, bind=engine)


# Query 1: List the disease code and the description of diseases that are caused by bacteria (pathogen) and were discovered before 1990.
print("Query 1: List the disease code and the description of diseases that are caused by bacteria (pathogen) and were discovered before 1990.")
with session() as conn:
    stmt = text(
        """
        SELECT 
            D.disease_code, D.description 
        FROM "disease" D
        JOIN "discovery" S ON D.disease_code = S.disease_code
        WHERE 
            D.pathogen = :pathogen AND extract('year' from S.first_enc_date) < :year
        """
    )

    rows = conn.execute(stmt, {"pathogen": "bacteria", "year": 1990}).fetchall()

    for row in rows:
        print(row)


print()

# Query 2: List the name, surname and degree of doctors who are not specialized in infectious diseases.
print("Query 2: List the name, surname and degree of doctors who are not specialized in infectious diseases.")
with session() as conn:
    stmt = text(
        """
        SELECT
            U.name, U.surname, D.degree
        FROM "doctor" D
        LEFT JOIN "users" U ON U.email = D.email
        WHERE U.email NOT IN (
            SELECT
            S.email
            FROM "specialize" S
            LEFT JOIN "disease_type" DT ON DT.id = S.id
            WHERE DT.description = :description
        );
        """
    )
    rows = conn.execute(stmt, {"description": 'infectious diseases'}).fetchall()

    for row in rows:
        print(row)

print()

# Query 3: List the name, surname and degree of doctors who are specialized in more than 2 disease types.
print("Query 3: List the name, surname and degree of doctors who are specialized in more than 2 disease types.")
with session() as conn:
    stmt = text(
        """
        SELECT
            U.name, U.surname, D.degree
        FROM "doctor" D
        LEFT JOIN "users" U ON D.email = U.email
        LEFT JOIN specialize S ON D.email = S.email
        GROUP BY U.name, U.surname, D.degree
        HAVING COUNT(*) > :count
        """
    )

    rows = conn.execute(stmt, {"count": 2}).fetchall()

    for row in rows:
        print(row)

print()

# Query 4: For each country list the cname and average salary of doctors who are specialized in virology.
print("Query 4: For each country list the cname and average salary of doctors who are specialized in virology.")
with session() as conn:
    stmt = text(
        """
        SELECT
            C.cname, AVG(U.salary)
        FROM doctor D
        LEFT JOIN specialize S ON D.email = S.email
        LEFT JOIN disease_type T ON S.id = T.id
        LEFT JOIN users U on D.email = U.email
        LEFT JOIN country C ON U.cname = C.cname
        WHERE
                T.description = :description
        GROUP BY C.cname;
        """
    )

    result = conn.execute(stmt, {"description":'virology'}).fetchall()

    for row in result:
        print(row)

print()


# Query 5: List the departments of public servants who report covid-19 cases in more than one country and the number of such public servants who work in these departments. (i.e Dept1 3 means that in the Dept1 department there are 3 such employees.)
print("Query 5: List the departments of public servants who report covid-19 cases in more than one country and the number of such public servants who work in these departments. (i.e Dept1 3 means that in the Dept1 department there are 3 such employees.)")
with session() as conn:
    stmt = text(
        """
        SELECT P.department, COUNT(DISTINCT P.email) AS count
        FROM public_servant P
        WHERE P.department IN
            (
                SELECT PS.department
                FROM public_servant PS
                LEFT JOIN record R on PS.email = r.email
                WHERE R.disease_code = :disease_code
                GROUP BY PS.department
                HAVING COUNT(*) > :count
            )
        GROUP BY P.department;
        """
    )

    result = conn.execute(stmt, {"disease_code": 'covid-19', "count": 1}).fetchall()

    for row in result:
        print(row)

print()

# Query 6: Double the salary of public servants who have recorded covid-19 patients more than 3 times.
print("Query 6: Double the salary of public servants who have recorded covid-19 patients more than 3 times.")
with session() as conn:
    stmt = text(
        """
        UPDATE users U
        SET salary = salary * 2
        FROM public_servant P
        JOIN record R ON P.email = R.email
        WHERE U.email IN (
            SELECT DISTINCT PS.email
            FROM public_servant PS
            LEFT JOIN record R on R.email = PS.email
            WHERE R.disease_code = :disease_code
            GROUP BY PS.email
            HAVING COUNT(*) > :count
        );
        """
    )

    r = conn.execute(stmt, {"disease_code":'covid-19', "count":3})
    conn.execute("COMMIT")

    print("Updated rows: ", r.rowcount)
    
    r.close()


print()
# Query 7: Delete the users whose name contain the substring “bek” or “gul” (e.g. Alibek, Gulsim).
print("Query 7: Delete the users whose name contain the substring “bek” or “gul” (e.g. Alibek, Gulsim).")
with session() as conn:
    stmt = text(
        """
        DELETE FROM users U
        WHERE 
            LOWER(U.name) LIKE :first OR LOWER(U.name) LIKE :second
        """
    )

    r = conn.execute(stmt, {"first": '%bek%', "second": '%gul%'})
    conn.execute("COMMIT")

    print("Deleted rows: ", r.rowcount)

print()
# Query 8: Create an index namely idx_pathogen on the pathogen field.
print("Query 8: Create an index namely idx_pathogen on the pathogen field.")
with session() as conn:
    stmt = text(
        """
        CREATE INDEX idx_pathogen ON disease (pathogen)
        """
    )

    r = conn.execute(stmt)
    conn.execute("COMMIT")


print()

# Query 9: List the email, name, and department of public servants who have created records where the number of patients is between 100000 and 999999
print("Query 9: List the email, name, and department of public servants who have created records where the number of patients is between 100000 and 999999")
with session() as conn:
    stmt = text(
        """
        SELECT DISTINCT
            P.email, U.name, P.department
        FROM public_servant P
        LEFT JOIN record R ON P.email = R.email
        LEFT JOIN users U ON P.email = U.email
        WHERE R.total_patients BETWEEN :first AND :second
        """
    )

    rows = conn.execute(stmt, {"first": 100000, "second": 999999}).fetchall()

    for row in rows:
        print(row)

print()

# Query 10: List the top 5 counties with the highest number of total patients recorded.
print("Query 10: List the top 5 counties with the highest number of total patients recorded.")
with session() as conn:
    stmt = text(
        """
        SELECT 
            C.cname
        FROM country C
        JOIN record R ON C.cname = R.cname
        GROUP BY C.cname
        ORDER BY SUM(R.total_patients) DESC
        LIMIT 5
        """
    )

    rows = conn.execute(stmt).fetchall()

    for row in rows:
        print(row)


print()

# Query 11: Group the diseases by disease type and the total number of patients treated.
print("Query 11: Group the diseases by disease type and the total number of patients treated.")
with session() as conn:
    stmt = text(
        """
        SELECT DISTINCT 
            T.description, COALESCE(SUM(R.total_patients), 0) AS total_patients
        FROM disease_type T
        LEFT JOIN disease D ON T.id = D.id
        LEFT JOIN record R ON D.disease_code = R.disease_code
        GROUP BY T.description
        """
    )

    rows = conn.execute(stmt).all()

    for row in rows:
        print(row)