ALTER TABLE public.persons ADD lastname varchar(255) NULL;
ALTER TABLE public.persons ADD gender char DEFAULT 'm' NOT NULL;
ALTER TABLE public.persons ADD CONSTRAINT check_allowed_genders CHECK (gender IN ('m', 'f'));

