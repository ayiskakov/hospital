--
-- PostgreSQL database dump
--

-- Dumped from database version 14.6
-- Dumped by pg_dump version 14.6

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: country; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.country (
    cname character varying(50) NOT NULL,
    population bigint NOT NULL
);


ALTER TABLE public.country OWNER TO postgres;

--
-- Name: discover; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.discover (
    cname character varying(50) NOT NULL,
    disease_code character varying(50) NOT NULL,
    first_enc_date date NOT NULL
);


ALTER TABLE public.discover OWNER TO postgres;

--
-- Name: disease; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.disease (
    id bigint NOT NULL,
    description character varying(140) NOT NULL,
    disease_code character varying(50) NOT NULL,
    pathogen character varying(20) NOT NULL
);


ALTER TABLE public.disease OWNER TO postgres;

--
-- Name: disease_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.disease_type (
    id bigint NOT NULL,
    description character varying(140) NOT NULL
);


ALTER TABLE public.disease_type OWNER TO postgres;

--
-- Name: disease_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.disease_type_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.disease_type_id_seq OWNER TO postgres;

--
-- Name: disease_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.disease_type_id_seq OWNED BY public.disease_type.id;


--
-- Name: doctor; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.doctor (
    email character varying(60) NOT NULL,
    degree character varying(20) NOT NULL
);


ALTER TABLE public.doctor OWNER TO postgres;

--
-- Name: public_servant; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.public_servant (
    email character varying(60) NOT NULL,
    department character varying(50) NOT NULL
);


ALTER TABLE public.public_servant OWNER TO postgres;

--
-- Name: record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.record (
    email character varying(60) NOT NULL,
    cname character varying(50) NOT NULL,
    disease_code character varying(50) NOT NULL,
    total_deaths bigint NOT NULL,
    total_patients bigint NOT NULL
);


ALTER TABLE public.record OWNER TO postgres;

--
-- Name: specialize; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.specialize (
    id bigint NOT NULL,
    email character varying(60) NOT NULL
);


ALTER TABLE public.specialize OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    email character varying(60) NOT NULL,
    name character varying(30) NOT NULL,
    surname character varying(40) NOT NULL,
    salary integer NOT NULL,
    phone character varying(20) NOT NULL,
    cname character varying(50) NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: disease_type id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.disease_type ALTER COLUMN id SET DEFAULT nextval('public.disease_type_id_seq'::regclass);


--
-- Data for Name: country; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.country (cname, population) FROM stdin;
Kazakhstan	19295600
China	1412600000
India	1375586000
Japan	125927902
Germany	84079811
France	67939000
United Kingdom	67886011
Italy	60483973
Russia	146745098
United States	329064917
\.


--
-- Data for Name: discover; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.discover (cname, disease_code, first_enc_date) FROM stdin;
Kazakhstan	G10.0	1989-05-09
China	R50	2000-03-02
India	GG1	1976-12-04
Japan	B15	2019-01-01
Russia	B15	2019-01-01
United States	B15	2019-01-01
Germany	B15	2019-01-01
United Kingdom	B15	2019-01-01
France	B15	2019-01-01
Italy	B15	2019-01-01
\.


--
-- Data for Name: disease; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.disease (id, description, disease_code, pathogen) FROM stdin;
1	Hepatitis A	B15	HAV
1	Hepatitis B	B180	HBV
1	covid-19	covid-19	SARS-CoV-2
1	Bacterial poisoning	GG1	bacteria
2	Scurvy	D53.2	C deficiency
2	Rickets	E55.0	D deficiency
2	Beriberi	E51.1	B1 deficiency
2	Pellagra	E52.0	bacteria
3	Cystic fibrosis	E84.0	CFTR mutation
3	Sickle cell anemia	D57.1	HBB mutation
3	Tay-Sachs disease	E75.02	HEXA mutation
3	Huntington's disease	G10.0	bacteria
4	Hypertension	I10	Unknown
4	Diabetes	E11	Unknown
4	Asthma	J45	Unknown
4	Allergy	R50	bacteria
\.


--
-- Data for Name: disease_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.disease_type (id, description) FROM stdin;
1	infectious diseases
2	deficiency diseases
3	hereditary diseases
4	physiological diseases
5	mental diseases
6	virology
7	congenital malformations
8	symptoms, signs, and ill-defined conditions
9	injury and poisoning
10	external causes of morbidity and mortality
\.


--
-- Data for Name: doctor; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.doctor (email, degree) FROM stdin;
bekbolat@hospital.com	QQ
gulmira@hospital.com	WW
gulsim@hospital.com	RR
ahobe@hospital.com	TT
karakat@hospital.com	SS
yelshi@hospital.com	QQ
kito@hospital.com	RR
louvre@hospital.com	SS
john@hospital.com	MD
json@hospital.com	SS
\.


--
-- Data for Name: public_servant; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.public_servant (email, department) FROM stdin;
aibek.tursanov@hospital.com	Dept1
saltanat@hospital.com	Dept1
meirbek@hospital.com	Dept1
nurlan@hospital.com	Dept2
erbolat@hospital.com	Dept3
kito@hospital.com	Dept2
shomala@hospital.com	Dept4
kirk@hospital.com	Dept5
ivan@hospital.com	Dept4
kemal@hospital.com	Dept3
\.


--
-- Data for Name: record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.record (email, cname, disease_code, total_deaths, total_patients) FROM stdin;
aibek.tursanov@hospital.com	Kazakhstan	G10.0	2342	34223
saltanat@hospital.com	United States	B15	32413	190000
aibek.tursanov@hospital.com	China	covid-19	43434	321312
aibek.tursanov@hospital.com	United States	covid-19	42300	33000
erbolat@hospital.com	Germany	covid-19	5653	432000
erbolat@hospital.com	Italy	covid-19	4333	402100
erbolat@hospital.com	Russia	covid-19	2334	487000
erbolat@hospital.com	United Kingdom	covid-19	300	23784
kirk@hospital.com	India	B15	5345	545445
kirk@hospital.com	India	E75.02	545	32111
\.


--
-- Data for Name: specialize; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.specialize (id, email) FROM stdin;
6	yelshi@hospital.com
6	john@hospital.com
6	kito@hospital.com
6	bekbolat@hospital.com
6	louvre@hospital.com
3	bekbolat@hospital.com
7	bekbolat@hospital.com
1	john@hospital.com
1	gulsim@hospital.com
7	karakat@hospital.com
8	karakat@hospital.com
6	ahobe@hospital.com
5	ahobe@hospital.com
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (email, name, surname, salary, phone, cname) FROM stdin;
aibek.tursanov@hospital.com	Aibek	Tursanov	1000	123456789	United Kingdom
bekbolat@hospital.com	Bekbolat	Kerey	1000	323456789	Kazakhstan
gulmira@hospital.com	Gulmira	Auezhay	1000	423456789	China
gulsim@hospital.com	Gulsim	Bektas	1000	723456789	India
ahobe@hospital.com	Ahobe	Koasa	1000	993456789	Japan
kito@hospital.com	Kunai	Jeko	4300	123456789	Japan
erbolat@hospital.com	Erbolat	Yerlanov	1000	188456789	Germany
john@hospital.com	John	Joe	1000	156756789	France
louvre@hospital.com	De	Paris	6500	345456789	France
nurlan@hospital.com	Nurlan	Karimov	1000	127656789	Italy
meirbek@hospital.com	Meirbek	Razorenov	1000	2347556789	Russia
saltanat@hospital.com	Saltanat	Neikolaeva	1000	4623456789	United States
karakat@hospital.com	Karakat	Danen	2300	6423456789	Kazakhstan
shomala@hospital.com	Shomala	Kerey	2300	123456789	Kazakhstan
yelshi@hospital.com	Yelshi	Kino	3900	123456789	Kazakhstan
kirk@hospital.com	Clark	Ken	1023	123523132	Russia
ivan@hospital.com	Ivan	Jperq	40233	99523132	Russia
kemal@hospital.com	Kemal	Jamal	40233	99523132	India
json@hospital.com	Geki	Qweras	33332	123213423	Japan
\.


--
-- Name: disease_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.disease_type_id_seq', 10, true);


--
-- Name: country country_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.country
    ADD CONSTRAINT country_pkey PRIMARY KEY (cname);


--
-- Name: discover discover_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.discover
    ADD CONSTRAINT discover_pkey PRIMARY KEY (cname, disease_code);


--
-- Name: disease disease_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.disease
    ADD CONSTRAINT disease_pkey PRIMARY KEY (disease_code);


--
-- Name: disease_type disease_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.disease_type
    ADD CONSTRAINT disease_type_pkey PRIMARY KEY (id);


--
-- Name: doctor doctor_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.doctor
    ADD CONSTRAINT doctor_pkey PRIMARY KEY (email);


--
-- Name: public_servant public_servant_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.public_servant
    ADD CONSTRAINT public_servant_pkey PRIMARY KEY (email);


--
-- Name: record record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.record
    ADD CONSTRAINT record_pkey PRIMARY KEY (email, cname, disease_code);


--
-- Name: specialize specialize_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.specialize
    ADD CONSTRAINT specialize_pkey PRIMARY KEY (id, email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (email);


--
-- Name: discover discover_cname_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.discover
    ADD CONSTRAINT discover_cname_fkey FOREIGN KEY (cname) REFERENCES public.country(cname) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: discover discover_disease_code_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.discover
    ADD CONSTRAINT discover_disease_code_fkey FOREIGN KEY (disease_code) REFERENCES public.disease(disease_code) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: disease disease_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.disease
    ADD CONSTRAINT disease_id_fkey FOREIGN KEY (id) REFERENCES public.disease_type(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: doctor doctor_email_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.doctor
    ADD CONSTRAINT doctor_email_fkey FOREIGN KEY (email) REFERENCES public.users(email) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: public_servant public_servant_email_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.public_servant
    ADD CONSTRAINT public_servant_email_fkey FOREIGN KEY (email) REFERENCES public.users(email) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: record record_cname_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.record
    ADD CONSTRAINT record_cname_fkey FOREIGN KEY (cname) REFERENCES public.country(cname) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: record record_disease_code_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.record
    ADD CONSTRAINT record_disease_code_fkey FOREIGN KEY (disease_code) REFERENCES public.disease(disease_code) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: record record_email_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.record
    ADD CONSTRAINT record_email_fkey FOREIGN KEY (email) REFERENCES public.public_servant(email) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: specialize specialize_email_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.specialize
    ADD CONSTRAINT specialize_email_fkey FOREIGN KEY (email) REFERENCES public.doctor(email) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: specialize specialize_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.specialize
    ADD CONSTRAINT specialize_id_fkey FOREIGN KEY (id) REFERENCES public.disease_type(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: users users_cname_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_cname_fkey FOREIGN KEY (cname) REFERENCES public.country(cname) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

