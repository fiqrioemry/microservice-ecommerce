import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

export function useFormSchema({ state, schema, onSubmit }) {
  const form = useForm({
    resolver: zodResolver(schema),
    defaultValues: state,
    mode: "onChange",
  });

  const handleSubmit = form.handleSubmit(onSubmit());

  return { form, handleSubmit };
}
