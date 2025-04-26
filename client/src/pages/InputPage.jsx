import React from "react";
import { addressSchema } from "@/lib/schema";
import { Button } from "@/components/ui/button";
import { addressState } from "@/components/formulir";
import { useForm, FormProvider } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useProfileStore } from "@/store/useProfileStore";
import LocationSelection from "@/components/beta/LocationSelection";
import CategorySelection from "../components/beta/CategorySelection";
import { toast } from "sonner";

const InputPage = () => {
  const { addAddress, loading } = useProfileStore();

  const methods = useForm({
    defaultValues: addressState,
    resolver: zodResolver(addressSchema),
  });

  const onSubmit = (data) => {
    console.log(data);
  };

  return (
    <FormProvider {...methods}>
      <form onSubmit={methods.handleSubmit(onSubmit)}>
        <CategorySelection />
        <LocationSelection />
        <Button
          type="submit"
          className="w-full"
          disabled={loading}
          isLoading={loading}
        />
      </form>
    </FormProvider>
  );
};

export default InputPage;
