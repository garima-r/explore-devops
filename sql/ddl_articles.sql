--
-- PostgreSQL database dump
--

--CREATE DATABASE "blog_demo";
\c blog_demo

CREATE TABLE articles(                                                                                                                                                                         
    id serial,                                                                                                                                                                                                  
    author varchar(50) NOT NULL,                                                                                                                                                                                
    title varchar(200) NOT NULL,                                                                                                                                                                                
    content text NOT NULL,
    PRIMARY KEY(author, title)
);

--
-- PostgreSQL database dump complete
--