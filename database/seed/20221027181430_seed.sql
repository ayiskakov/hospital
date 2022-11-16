-- +goose Up
-- +goose StatementBegin


INSERT INTO "public.disease_type"
    ("description")
VALUES
    ('infectious diseases'),
    ('deficiency diseases'),
    ('hereditary diseases'),
    ('physiological diseases'),
    ('mental diseases'),
    ('virology'),
    ('congenital malformations'),
    ('symptoms, signs, and ill-defined conditions'),
    ('injury and poisoning'),
    ('external causes of morbidity and mortality');

INSERT INTO "public.country"
    (cname, population)
VALUES
    ('Kazakhstan', 19295600),
    ('China', 1412600000),
    ('India', 1375586000),
    ('Japan', 125927902),
    ('Germany', 84079811),
    ('France', 67939000),
    ('United Kingdom', 67886011),
    ('Italy', 60483973),
    ('Russia', 146745098),
    ('United States', 329064917);


INSERT INTO "public.disease"
    ("id", "description", "disease_code", "pathogen")
VALUES
    (1, 'Hepatitis A', 'B15', 'HAV'),
    (1, 'Hepatitis B', 'B180', 'HBV'),
    (1, 'covid-19', 'U08', 'SARS-CoV-2'),
    (1, 'Bacterial poisoning', 'GG1', 'bacteria');

-- Deficiency diseases only
INSERT INTO "public.disease"
    ("id", "description", "disease_code", "pathogen")
VALUES
    (2, 'Scurvy', 'D53.2', 'C deficiency'),
    (2, 'Rickets', 'E55.0', 'D deficiency'),
    (2, 'Beriberi', 'E51.1', 'B1 deficiency'),
    (2, 'Pellagra', 'E52.0', 'bacteria');

-- Hereditary diseases only
INSERT INTO "public.disease"
    ("id", "description", "disease_code", "pathogen")
VALUES
    (3, 'Cystic fibrosis', 'E84.0', 'CFTR mutation'),
    (3, 'Sickle cell anemia', 'D57.1', 'HBB mutation'),
    (3, 'Tay-Sachs disease', 'E75.02', 'HEXA mutation'),
    (3, 'Huntington''s disease', 'G10.0', 'bacteria');

-- Physiological diseases only
INSERT INTO "public.disease"
    ("id", "description", "disease_code", "pathogen")
VALUES
    (4, 'Hypertension', 'I10', 'Unknown'),
    (4, 'Diabetes', 'E11', 'Unknown'),
    (4, 'Asthma', 'J45', 'Unknown'),
    (4, 'Allergy', 'R50', 'bacteria');


INSERT INTO "public.discovery"
    ("cname", "disease_code", "first_enc_date")
VALUES
    ('Kazakhstan', 'G10.0', '1989-05-09'),
    ('China', 'R50', '2000-03-02'),
    ('India', 'GG1', '1976-12-04'),
    ('Japan', 'B15', '2019-01-01'),
    ('Russia', 'B15', '2019-01-01'),
    ('United States', 'B15', '2019-01-01'),
    ('Germany', 'B15', '2019-01-01'),
    ('United Kingdom', 'B15', '2019-01-01'),
    ('France', 'B15', '2019-01-01'),
    ('Italy', 'B15', '2019-01-01');


INSERT INTO "public.user"
    ("email", "name", "surname", "salary", "phone", "cname")
VALUES
    ('aibek.tursanov@hospital.com', 'Aibek', 'Tursanov', 1000, '123456789', 'United Kingdom'),
    ('bekbolat@hospital.com', 'Bekbolat', 'Kerey', 1000, '123456789', 'Kazakhstan'),
    ('gulmira@hospital.com', 'Gulmira', 'Auezhay', 1000, '123456789', 'China'),
    ('gulsim@hospital.com', 'Gulsim', 'Bektas', 1000, '123456789', 'India'),
    ('ahobe@hospital.com', 'Ahobe', 'Koasa', 1000, '123456789', 'Japan'),
    ('erbolat@hospital.com', 'Erbolat', 'Yerlanov', 1000, '123456789', 'Germany'),
    ('john@hospital.com', 'John', 'Joe', 1000, '123456789', 'France'),
    ('nurlan@hospital.com', 'Nurlan', 'Karimov', 1000, '123456789', 'Italy'),
    ('meirbek@hospital.com', 'Meirbek', 'Razorenov', 1000, '123456789', 'Russia'),
    ('saltanat@hospital.com', 'Saltanat', 'Neikolaeva', 1000, '123456789', 'United States'),
    ('karakat@hospital.com', 'Karakat', 'Danen', 2300, '123456789', 'Kazakhstan'),
    ('shomala@hospital.com', 'Shomala', 'Kerey', 2300, '123456789', 'Kazakhstan');

INSERT INTO "public.public_servant"
    ("email", "department")
VALUES
    ('aibek.tursanov@hospital.com', 'Dept1'),
    ('saltanat@hospital.com', 'Dept1'),
    ('meirbek@hospital.com', 'Dept1'),
    ('nurlan@hospital.com', 'Dept2'),
    ('erbolat@hospital.com', 'Dept3'),
    ('shomala@hospital.com', 'Dept4');

INSERT INTO "public.doctor"
    ("email", "degree")
VALUES
    ('bekbolat@hospital.com', 'MD'),
    ('gulmira@hospital.com', 'DO'),
    ('gulsim@hospital.com', 'MD'),
    ('ahobe@hospital.com', 'MD'),
    ('karakat@hospital.com', 'MD'),
    ('john@hospital.com', 'MD');

INSERT INTO "public.specialize"
    ("id", "email")
VALUES
    (6, 'john@hospital.com'),
    (6, 'bekbolat@hospital.com'),
    (3, 'bekbolat@hospital.com'),
    (7, 'bekbolat@hospital.com'),
    (1, 'john@hospital.com'),
    (1, 'gulsim@hospital.com'),
    (7, 'karakat@hospital.com'),
    (8, 'karakat@hospital.com'),
    (6, 'ahobe@hospital.com'),
    (5, 'ahobe@hospital.com');


INSERT INTO "public.record"
    ("email", "cname", "disease_code", "total_deaths", "total_patients")
VALUES
    ('aibek.tursanov@hospital.com', 'Kazakhstan', 'G10.0', 100, 1000),
    ('saltanat@hospital.com', 'United States', 'B15', 200, 190000),
    ('aibek.tursanov@hospital.com', 'China', 'U08', 300, 3000),
    ('aibek.tursanov@hospital.com', 'United States', 'U08', 300, 3000),
    ('erbolat@hospital.com', 'Germany', 'U08', 400, 400000),
    ('erbolat@hospital.com', 'United Kingdom', 'U08', 300, 23784);






-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
