CREATE TABLE public.user
(
    user_id int NOT NULL PRIMARY KEY,
    user_name VARCHAR(100) NOT NULL
);

CREATE TABLE public.segment
(
    segment_id SERIAL PRIMARY KEY,
    segment_name VARCHAR(100) NOT NULL
);

CREATE TABLE public.users_in_segm
(
    user_id int ,
    segment_id int ,
    ttl DATE,
    CONSTRAINT id PRIMARY KEY (user_id, segment_id),
    FOREIGN KEY (user_id)  REFERENCES public.user(user_id),
    FOREIGN KEY (segment_id)  REFERENCES public.segment(segment_id)
);


CREATE TABLE public.statistics
(
    id SERIAL PRIMARY KEY,
    user_id int NOT NULL,
    segment_name VARCHAR(100) NOT NULL,
    date_event DATE,
    in_event bool,
    FOREIGN KEY (user_id)  REFERENCES public.user(user_id)
);

-- Вставка значений в таблицу public.user
-- INSERT INTO public.user (user_id, user_name)
-- VALUES
--     (1, 'User A'),
--     (2, 'User B'),
--     (3, 'User C'),
--     (4, 'User D'),
--     (5, 'User E'),
--     (6, 'User F'),
--     (7, 'User G');
--
--
-- INSERT INTO public.segment  (segment_name)
-- VALUES
--     ( 'Segment X'),
--     ( 'Segment Y'),
--     ( 'Segment Z');
