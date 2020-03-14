CREATE OR REPLACE FUNCTION trigger_set_updated_stamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON public.project
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_updated_stamp();
  
CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON public.project_information
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_updated_stamp();

CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON public.version
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_updated_stamp();