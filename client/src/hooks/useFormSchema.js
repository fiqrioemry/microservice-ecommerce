import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

export function useFormSchema({ schema, defaultValues, onSubmit }) {
  const form = useForm({
    resolver: zodResolver(schema),
    defaultValues,
    mode: "onChange",
  });

  const handleSubmit = form.handleSubmit(onSubmit);

  return { form, handleSubmit };
}
