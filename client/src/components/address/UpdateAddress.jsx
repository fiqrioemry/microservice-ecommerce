// src/components/address/DeleteAddress.jsx
import React from "react";
import { Edit2 } from "lucide-react";
import { addressSchema } from "@/lib/schema";
import { FormDialog } from "@/components/form/FormDialog";
import { useProfileStore } from "@/store/useProfileStore";
import { InputElement } from "@/components/input/InputElement";
import { SwitchElement } from "@/components/input/SwitchElement";
import { useProfileManagement } from "@/hooks/useProfileManagement";
import LocationSelection from "@/components/input/LocationSelection";

const UpdateAddress = ({ address }) => {
  const { updateAddress, loading } = useProfileStore();
  const { invalidateUserAddresses } = useProfileManagement();

  const handleUpdate = async (formData) => {
    const normalizedData = {
      ...formData,
      cityId: Number(formData.cityId),
      provinceId: Number(formData.provinceId),
      districtId: Number(formData.districtId),
      subdistrictId: Number(formData.subdistrictId),
      postalCodeId: Number(formData.postalCodeId),
    };
    await updateAddress(address.id, normalizedData);
    invalidateUserAddresses();
  };

  return (
    <FormDialog
      loading={loading}
      state={address}
      action={handleUpdate}
      schema={addressSchema}
      title="Update Address"
      buttonText={
        <button className="btn btn-secondary">
          <Edit2 className="text-white-500" size={18} />
        </button>
      }
    >
      <InputElement
        name="name"
        label="Nama Penerima"
        placeholder="Masukkan nama penerima"
      />
      <InputElement
        name="phone"
        isNumber={true}
        label="Nomor Telepon"
        placeholder="Masukkan nomor penerima"
      />
      <InputElement
        name="address"
        isTextArea
        label="Alamat Lengkap"
        placeholder="Masukkan alamat lengkap"
      />

      <LocationSelection />

      <SwitchElement name="isMain" label="Atur sebagai alamat utama?" />
    </FormDialog>
  );
};

export default UpdateAddress;
