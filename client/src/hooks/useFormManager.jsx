import { useEffect, useState } from "react";
import { useProductStore } from "../store/useProductStore";

export const useFormManager = (schema) => {
  const { loadOptions } = useProductStore();
  const [formSchema, setFormSchema] = useState(schema);

  useEffect(() => {
    const injectAsyncOptions = async () => {
      const updatedSchema = await Promise.all(
        schema.map(async (field) => {
          if (field.isAsync && field.loaderKey) {
            const options = await loadOptions(field.loaderKey);
            return { ...field, options };
          }
          return field;
        })
      );
      setFormSchema(updatedSchema);
    };

    injectAsyncOptions();
  }, []);

  const defaultValues = Object.fromEntries(
    schema.map((field) => [field.name, ""])
  );

  return { formSchema, defaultValues };
};
