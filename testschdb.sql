PGDMP             
            |        	   testschdb    15.6 (Debian 15.6-0+deb12u1)    15.6 (Debian 15.6-0+deb12u1)     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16558 	   testschdb    DATABASE     u   CREATE DATABASE testschdb WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_GB.UTF-8';
    DROP DATABASE testschdb;
                postgres    false            �           0    0    DATABASE testschdb    ACL     ,   GRANT ALL ON DATABASE testschdb TO testusr;
                   postgres    false    3478            �            1259    16559    students    TABLE     �   CREATE TABLE public.students (
    name character varying(100),
    klass character varying(30),
    grade character varying(30)
);
    DROP TABLE public.students;
       public         heap    postgres    false            �          0    16559    students 
   TABLE DATA           6   COPY public.students (name, klass, grade) FROM stdin;
    public          postgres    false    214          �   h   x�˱� �����1\
#���ta kcr��}�b�/�����0,��U��qp����p�-�6B���3ߕ̞kT�B�na8ҧԿ����=A��f>�?�     