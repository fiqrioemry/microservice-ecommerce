// src/components/address/AddAddress.jsx
import React from "react";
import { PlusCircle } from "lucide-react";
import { addressSchema } from "@/lib/schema";
import { addressState } from "@/lib/constant";
import { FormDialog } from "@/components/form/FormDialog";
import { useProfileStore } from "@/store/useProfileStore";
import { InputElement } from "@/components/input/InputElement";
import { SwitchElement } from "@/components/input/SwitchElement";
import { useProfileManagement } from "@/hooks/useProfileManagement";
import LocationSelection from "@/components/input/LocationSelection";

const AddAddress = () => {
  const { addAddress, loading } = useProfileStore();
  const { invalidateUserAddresses } = useProfileManagement();

  // const handleAddAddress = async (formData) => {
  //   const normalizedData = {
  //     ...formData,
  //     cityId: Number(formData.cityId),
  //     provinceId: Number(formData.provinceId),
  //     districtId: Number(formData.districtId),
  //     subdistrictId: Number(formData.subdistrictId),
  //     postalCodeId: Number(formData.postalCodeId),
  //   };
  //   await addAddress(normalizedData);
  //   invalidateUserAddresses();
  // };

  return (
    <FormDialog
      loading={loading}
      action={addAddress}
      state={addressState}
      schema={addressSchema}
      title="Add New Address"
      buttonText={
        <button className="btn btn-primary gap-4">
          <PlusCircle size={18} />
          <span>New Address</span>
        </button>
      }
    >
      <InputElement
        name="name"
        label="Nama penerima"
        placeholder="Masukkan nama penerima"
      />
      <InputElement
        name="phone"
        // isNumber={true}
        maxLength={12}
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
