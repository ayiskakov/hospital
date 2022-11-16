from sqlalchemy import create_engine, text

# Create a connection to the database
engine = create_engine(
    'postgresql://postgres:postgres@localhost:5432/postgres', 
    echo=True, future=True
)


# Query 1: List the disease code and the description of diseases that are caused by bacteria (pathogen) and were discovered before 1990.
print("Query 1: List the disease code and the description of diseases that are caused by bacteria (pathogen) and were discovered before 1990.")
with engine.connect() as conn:
    stmt = text(
        """
        SELECT 
            D.disease_code, D.description 
        FROM public.disease D
        JOIN public.discovery S ON D.disease_code = S.disease_code
        WHERE 
            D.pathogen = ? AND YEAR(S.first_enc_data) < ?
        """
    )

    rows = conn.execute(stmt, 'bacteria', 1990).all()

    for row in rows:
        print(row)

print()

# Query 2: List the name, surname and degree of doctors who are not specialized in infectious diseases.
print("Query 2: List the name, surname and degree of doctors who are not specialized in infectious diseases.")
with engine.connect() as conn:
    stmt = text(
        """
        SELECT 
            D.name, D.surname, D.degree 
        FROM public.doctor D
        JOIN public.specialize S ON D.email = S.email
        JOIN public.disease_type T ON S.id = T.id
        WHERE 
            T.description <> ?
        """
    )

    rows = conn.execute(stmt, 'infectious diseases').all()

    for row in rows:
        print(row)

print()

# Query 3: List the name, surname and degree of doctors who are specialized in more than 2 disease types.
print("Query 3: List the name, surname and degree of doctors who are specialized in more than 2 disease types.")
with engine.connect() as conn:
    stmt = text(
        """
        SELECT 
            D.name, D.surname, D.degree 
        FROM public.doctor D
        JOIN public.specialize S ON D.email = S.email
        GROUP BY D.name, D.surname, D.degree
        HAVING COUNT(*) > ?
        """
    )

    rows = conn.execute(stmt, 2).all()

    for row in rows:
        print(row)

print()

# Query 4: For each country list the cname and average salary of doctors who are specialized in virology.
print("Query 4: For each country list the cname and average salary of doctors who are specialized in virology.")
with engine.connect() as conn:
    stmt = text(
        """
        SELECT 
            C.cname, AVG(U.salary) 
        FROM public.doctor D
        JOIN public.specialize S ON D.email = S.email
        JOIN public.disease_type T ON S.id = T.id
        JOIN public.country C ON D.cname = C.cname
        JOIN public.user U on D.email = U.email
        WHERE 
            T.description = ?
        GROUP BY C.cname
        """
    )

    result = conn.execute(stmt, 'virology').all()

    for row in result:
        print(row)

print()

# Query 5: List the departments of public servants who report covid-19 cases in more than one country and the number of such public servants who work in these departments. (i.e Dept1 3 means that in the Dept1 department there are 3 such employees.)
print("Query 5: List the departments of public servants who report covid-19 cases in more than one country and the number of such public servants who work in these departments. (i.e Dept1 3 means that in the Dept1 department there are 3 such employees.)")
with engine.connect() as conn:
    stmt = text(
        """
        SELECT
            P.department, COUNT(*) 
        FROM public.public_servant P
        JOIN public.record R ON P.email = R.email
        WHERE 
            R.disease_code = ?
        GROUP BY P.department
        HAVING COUNT(*) > ?
        """
    )

    result = conn.execute(stmt, 'covid-19', 1).all()

    for row in result:
        print(row)

print()

# Query 6: Double the salary of public servants who have recorded covid-19 patients more than 3 times.
print("Query 6: Double the salary of public servants who have recorded covid-19 patients more than 3 times.")
with engine.connect() as conn:
    stmt = text(
        """
        UPDATE public.user U
        SET salary = salary * 2
        FROM public.public_servant P
        JOIN public.record R ON P.email = R.email
        WHERE 
            R.disease_code = ? AND P.email = U.email
        GROUP BY P.email
        HAVING COUNT(*) > ?
        """
    )

    conn.execute(stmt, 'covid-19', 3)


# Query 7: Delete the users whose name contain the substring “bek” or “gul” (e.g. Alibek, Gulsim).
print("Query 7: Delete the users whose name contain the substring “bek” or “gul” (e.g. Alibek, Gulsim).")
with engine.connect() as conn:
    stmt = text(
        """
        DELETE FROM public.user U
        WHERE 
            LOWER(U.name) LIKE ? OR LOWER(U.name) LIKE ?
        """
    )

    conn.execute(stmt, '%bek%', '%gul%')


# Query 8: Create an index namely idx_pathogen on the pathogen field.
print("Query 8: Create an index namely idx_pathogen on the pathogen field.")
with engine.connect() as conn:
    stmt = text(
        """
        CREATE INDEX idx_pathogen ON public.disease (pathogen)
        """
    )

    conn.execute(stmt)


# Query 9: List the email, name, and department of public servants who have created records where the number of patients is between 100000 and 999999
print("Query 9: List the email, name, and department of public servants who have created records where the number of patients is between 100000 and 999999")
with engine.connect() as conn:
    stmt = text(
        """
        SELECT 
            P.email, U.name, P.department 
        FROM public.public_servant P
        JOIN public.record R ON P.email = R.email
        JOIN public.user U ON P.email = U.email
        WHERE 
            R.total_patients BETWEEN ? AND ?
        """
    )

    rows = conn.execute(stmt, 100000, 999999).all()

    for row in rows:
        print(row)

print()

# Query 10: List the top 5 counties with the highest number of total patients recorded.
print("Query 10: List the top 5 counties with the highest number of total patients recorded.")
with engine.connect() as conn:
    stmt = text(
        """
        SELECT 
            C.cname, SUM(R.total_patients) AS total_patients
        FROM public.country C
        JOIN public.record R ON C.cname = R.cname
        GROUP BY C.cname
        ORDER BY total_patients DESC
        LIMIT 5
        """
    )

    rows = conn.execute(stmt).all()

    for row in rows:
        print(row)


print()

# Query 11: Group the diseases by disease type and the total number of patients treated.
print("Query 11: Group the diseases by disease type and the total number of patients treated.")
with engine.connect() as conn:
    stmt = text(
        """
        SELECT 
            T.description, SUM(R.total_patients) AS total_patients
        FROM public.disease_type T
        JOIN public.disease D ON T.id = D.id
        JOIN public.record R ON D.disease_code = R.disease_code
        GROUP BY T.description
        """
    )

    rows = conn.execute(stmt).all()

    for row in rows:
        print(row)