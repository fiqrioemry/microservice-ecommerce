import { FormGenerator } from "./FormGenerator";
import { Button } from "@/components/ui/button";
import { useFormSchema } from "@/hooks/useFormSchema";
import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/dialog";

export function DialogForm({
  schema,
  defaultValues,
  fields,
  action,
  buttonLabel,
  title,
  loading,
}) {
  const { form, handleSubmit } = useFormSchema({
    schema,
    defaultValues,
    onSubmit: action,
  });

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>{buttonLabel}</Button>
      </DialogTrigger>
      <DialogContent>
        <h4 className="text-lg font-bold mb-4">{title}</h4>
        <form onSubmit={handleSubmit} className="space-y-4">
          <FormGenerator control={form.control} fields={fields} />
          <Button type="submit" disabled={!form.formState.isValid || loading}>
            {loading ? "Loading..." : "Submit"}
          </Button>
        </form>
      </DialogContent>
    </Dialog>
  );
}
