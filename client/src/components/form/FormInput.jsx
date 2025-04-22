"use client";

import { useForm } from "react-hook-form";
import { Button } from "@/components/ui/button";
import { zodResolver } from "@hookform/resolvers/zod";
import { FormRenderer } from "@/components/form/FormRenderer";
import { useFormManager } from "@/hooks/useFormManager";

const FormInput = ({ action, schema, validation, submitText }) => {
  const { formSchema, defaultValues } = useFormManager(schema);
  const {
    handleSubmit,
    control,
    formState: { errors },
  } = useForm({
    defaultValues,
    resolver: zodResolver(validation),
  });

  const onSubmit = (data) => {
    action(data);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <div className="grid grid-cols-4 gap-4 mb-4">
        <FormRenderer control={control} schema={formSchema} />
      </div>
      <Button type="submit">{submitText}</Button>
    </form>
  );
};
export default FormInput;
