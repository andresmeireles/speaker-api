-- Revert adding the lastname column
ALTER TABLE public.persons DROP COLUMN IF EXISTS lastname;

-- Revert adding the gender column and the constraint
ALTER TABLE public.persons DROP COLUMN IF EXISTS gender;
ALTER TABLE public.persons DROP CONSTRAINT IF EXISTS check_allowed_genders;

