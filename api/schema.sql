--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2 (Debian 17.2-1.pgdg120+1)
-- Dumped by pg_dump version 17.2 (Ubuntu 17.2-1.pgdg22.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: feed_tag; Type: TABLE; Schema: public; Owner: iacnaylor
--

CREATE TABLE public.feed_tag (
    feed_id integer NOT NULL,
    tag_id integer NOT NULL
);


ALTER TABLE public.feed_tag OWNER TO iacnaylor;

--
-- Name: news_feed; Type: TABLE; Schema: public; Owner: iacnaylor
--

CREATE TABLE public.news_feed (
    id integer NOT NULL,
    feed_author text NOT NULL,
    feed_url text NOT NULL,
    added_by integer NOT NULL
);


ALTER TABLE public.news_feed OWNER TO iacnaylor;

--
-- Name: news_feed_id_seq; Type: SEQUENCE; Schema: public; Owner: iacnaylor
--

ALTER TABLE public.news_feed ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.news_feed_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: tag; Type: TABLE; Schema: public; Owner: iacnaylor
--

CREATE TABLE public.tag (
    id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.tag OWNER TO iacnaylor;

--
-- Name: tag_id_seq; Type: SEQUENCE; Schema: public; Owner: iacnaylor
--

ALTER TABLE public.tag ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.tag_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: user; Type: TABLE; Schema: public; Owner: iacnaylor
--

CREATE TABLE public."user" (
    id integer NOT NULL,
    username text NOT NULL,
    password text NOT NULL
);


ALTER TABLE public."user" OWNER TO iacnaylor;

--
-- Name: user_feed; Type: TABLE; Schema: public; Owner: iacnaylor
--

CREATE TABLE public.user_feed (
    user_id integer NOT NULL,
    feed_id integer NOT NULL
);


ALTER TABLE public.user_feed OWNER TO iacnaylor;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: iacnaylor
--

ALTER TABLE public."user" ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: feed_tag PK_FeedID_TagID; Type: CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.feed_tag
    ADD CONSTRAINT "PK_FeedID_TagID" PRIMARY KEY (feed_id, tag_id);


--
-- Name: news_feed PK_FeedId; Type: CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.news_feed
    ADD CONSTRAINT "PK_FeedId" PRIMARY KEY (id);


--
-- Name: user_feed PK_UserID_FeedID; Type: CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.user_feed
    ADD CONSTRAINT "PK_UserID_FeedID" PRIMARY KEY (user_id, feed_id);


--
-- Name: tag UQ_Tag; Type: CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT "UQ_Tag" UNIQUE (name);


--
-- Name: user UQ_Username; Type: CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT "UQ_Username" UNIQUE (username);


--
-- Name: tag tag_pkey; Type: CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_pkey PRIMARY KEY (id);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: news_feed FK_AddedBy_UserId; Type: FK CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.news_feed
    ADD CONSTRAINT "FK_AddedBy_UserId" FOREIGN KEY (added_by) REFERENCES public."user"(id);


--
-- Name: user_feed FK_FeedId_Feed; Type: FK CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.user_feed
    ADD CONSTRAINT "FK_FeedId_Feed" FOREIGN KEY (feed_id) REFERENCES public.news_feed(id) ON DELETE CASCADE;


--
-- Name: feed_tag FK_FeedId_Feed; Type: FK CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.feed_tag
    ADD CONSTRAINT "FK_FeedId_Feed" FOREIGN KEY (feed_id) REFERENCES public.news_feed(id) ON DELETE CASCADE;


--
-- Name: feed_tag FK_TagId_Tag; Type: FK CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.feed_tag
    ADD CONSTRAINT "FK_TagId_Tag" FOREIGN KEY (tag_id) REFERENCES public.tag(id) ON DELETE CASCADE;


--
-- Name: user_feed FK_UserId_User; Type: FK CONSTRAINT; Schema: public; Owner: iacnaylor
--

ALTER TABLE ONLY public.user_feed
    ADD CONSTRAINT "FK_UserId_User" FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

