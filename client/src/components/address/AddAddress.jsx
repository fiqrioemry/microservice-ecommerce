import React from "react";
import { addressSchema } from "@/lib/schema";
import { addressState } from "@/lib/constant";
import { FormDialog } from "@/components/form/FormDialog";
import { useProfileStore } from "@/store/useProfileStore";
import { InputElement } from "@/components/input/InputElement";
import { SwitchElement } from "@/components/input/SwitchElement";
import LocationSelection from "@/components/input/LocationSelection";

const AddAddress = () => {
  const { addAddress, loading } = useProfileStore();
  return (
    <FormDialog
      loading={loading}
      action={addAddress}
      state={addressState}
      schema={addressSchema}
      title="Add New Address"
      buttonText="New Address"
    >
      <InputElement
        name="name"
        label="Nama penerima"
        placeholder="Masukkan nama penerima"
      />
      <InputElement
        name="phone"
        isNumber={true}
        label="Nomor Telepon"
        placeholder="Masukkan Nomor Penerima"
      />
      <InputElement
        name="address"
        label="Alamat Penerima"
        isTextArea={true}
        placeholder="Masukkan Alamat Penerima"
      />

      <LocationSelection />

      <SwitchElement name="isMain" label="Atur sebagai alamat utama ?" />
    </FormDialog>
  );
};

export default AddAddress;
