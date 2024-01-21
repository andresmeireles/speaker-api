ALTER TABLE invites ADD status smallint NOT NULL DEFAULT 0;
ALTER TABLE public.invites ADD CONSTRAINT ck_status CHECK (((status >= 0) AND (status <= 6)));
