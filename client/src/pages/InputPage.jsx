import React from "react";
import { addressSchema } from "@/lib/schema";
import { Button } from "@/components/ui/button";
import { addressState } from "@/components/formulir";
import { useForm, FormProvider } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useProfileStore } from "@/store/useProfileStore";
import LocationSelection from "@/components/beta/LocationSelection";

const InputPage = () => {
  const { addAddress, loading } = useProfileStore();

  const methods = useForm({
    defaultValues: addressState,
    resolver: zodResolver(addressSchema),
  });

  const onSubmit = (data) => {
    addAddress(data);
  };

  return (
    <FormProvider {...methods}>
      <form onSubmit={methods.handleSubmit(onSubmit)}>
        <LocationSelection />
        <Button
          variant="primary"
          className="w-full"
          isLoading={loading}
          disabled={loading}
        >
          Submit New Location
        </Button>
      </form>
    </FormProvider>
  );
};

export default InputPage;
