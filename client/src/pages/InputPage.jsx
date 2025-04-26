import React from "react";
import { addressSchema } from "@/lib/schema";
import { addressState } from "@/components/formulir";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, FormProvider } from "react-hook-form";
import { useProfileStore } from "@/store/useProfileStore";
import { SubmitButton } from "@/components/beta/SubmitButton";
import { InputElement } from "@/components/beta/InputElement";
import { SwitchElement } from "@/components/beta/SwitchElement";
import LocationSelection from "@/components/beta/LocationSelection";
import { SelectElement } from "../components/beta/SelectElement";
import { useCategoriesQuery } from "../hooks/useCategoriesQuery";

const InputPage = () => {
  const { addAddress, loading } = useProfileStore();

  const methods = useForm({
    defaultValues: addressState,
    resolver: zodResolver(addressSchema),
  });

  const { data = {} } = useCategoriesQuery();
  const categories = data.categories || [];

  const onSubmit = async (data) => {
    await addAddress(data);
  };

  return (
    <FormProvider {...methods}>
      <form onSubmit={methods.handleSubmit(onSubmit)} className="max-w-2xl">
        <InputElement
          name="name"
          label="Nama Penerima"
          placeholder="Masukkan nama Penerima"
        />
        <SelectElement
          name="categoryId"
          label="Category"
          placeholder="Select Category "
          options={categories} // langsung, tanpa map
        />

        <SelectElement
          name="subcategoryId"
          label="Subcategory"
          placeholder="Select Subcategory "
          options={subcategories} // langsung, tanpa map
        />
        <InputElement
          name="address"
          label="Alamat Lengkap"
          isTextArea={true}
          placeholder="Masukkan Alamat Penerima"
        />
        <LocationSelection />

        <InputElement
          name="phone"
          label="Nomor Penerima"
          placeholder="Masukkan Nomor Penerima"
        />
        <SwitchElement name="isMain" label="Atur alamat utama" />
        <SubmitButton
          isLoading={loading}
          disabled={!methods.formState.isValid}
        />
      </form>
    </FormProvider>
  );
};

export default InputPage;
