// src/components/form/FormInput.jsx
import React from "react";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, FormProvider } from "react-hook-form";

const FormInput = ({ action, state, schema, children }) => {
  const methods = useForm({
    defaultValues: state,
    resolver: zodResolver(schema),
  });

  const onSubmit = async (data) => {
    await action(data);
  };

  return (
    <FormProvider {...methods}>
      <form
        onSubmit={methods.handleSubmit(onSubmit)}
        className="grid-cols-2 space-y-2"
      >
        {typeof children === "function" ? children(methods) : children}
      </form>
    </FormProvider>
  );
};

export { FormInput };
