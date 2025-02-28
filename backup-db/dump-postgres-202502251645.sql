PGDMP      -                }            postgres    16.8    16.8     2           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            3           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            4           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            5           1262    5    postgres    DATABASE     n   CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en-US';
    DROP DATABASE postgres;
                postgres    false            6           0    0    DATABASE postgres    COMMENT     N   COMMENT ON DATABASE postgres IS 'default administrative connection database';
                   postgres    false    4917                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                pg_database_owner    false            7           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                   pg_database_owner    false    6            �            1259    16430    to_do_groups    TABLE     �   CREATE TABLE public.to_do_groups (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    group_name character varying(255) NOT NULL
);
     DROP TABLE public.to_do_groups;
       public         heap    postgres    false    6    6    6            �            1259    16421    to_do_lists    TABLE     �   CREATE TABLE public.to_do_lists (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    to_do_task text,
    status boolean DEFAULT false NOT NULL,
    group_id character varying NOT NULL,
    owner_id character varying NOT NULL
);
    DROP TABLE public.to_do_lists;
       public         heap    postgres    false    6    6    6            �            1259    16411    users    TABLE     �   CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(255) NOT NULL,
    username character varying(255) NOT NULL,
    password text NOT NULL
);
    DROP TABLE public.users;
       public         heap    postgres    false    6    6    6            /          0    16430    to_do_groups 
   TABLE DATA           6   COPY public.to_do_groups (id, group_name) FROM stdin;
    public          postgres    false    219   '       .          0    16421    to_do_lists 
   TABLE DATA           Q   COPY public.to_do_lists (id, to_do_task, status, group_id, owner_id) FROM stdin;
    public          postgres    false    218   D       -          0    16411    users 
   TABLE DATA           =   COPY public.users (id, name, username, password) FROM stdin;
    public          postgres    false    217   a       �           2606    16435    to_do_groups to_do_groups_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.to_do_groups
    ADD CONSTRAINT to_do_groups_pkey PRIMARY KEY (id);
 H   ALTER TABLE ONLY public.to_do_groups DROP CONSTRAINT to_do_groups_pkey;
       public            postgres    false    219            �           2606    16429    to_do_lists to_do_lists_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.to_do_lists
    ADD CONSTRAINT to_do_lists_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.to_do_lists DROP CONSTRAINT to_do_lists_pkey;
       public            postgres    false    218            �           2606    16418    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    217            �           2606    16420    users users_username_key 
   CONSTRAINT     W   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);
 B   ALTER TABLE ONLY public.users DROP CONSTRAINT users_username_key;
       public            postgres    false    217            /      x������ � �      .      x������ � �      -   i   x�=�1�0 �9��Q��q�����[� "�G��r�U�h��PjTzou��E�M�x���[���>�MEi�
��X;��8����,�*�̤��y��X"�     