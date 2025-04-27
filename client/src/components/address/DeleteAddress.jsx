/* eslint-disable react/prop-types */
// src/components/address/DeleteAddress.jsx
import { useProfileStore } from "@/store/useProfileStore";
import { FormDelete } from "@/components/form/FormDelete";
import { useProfileManagement } from "@/hooks/useProfileManagement";

const DeleteAddress = ({ address }) => {
  const { deleteAddress, loading } = useProfileStore();
  const { invalidateUserAddresses } = useProfileManagement();

  const handleDelete = async () => {
    const normalizedData = {
      ...formData,
      cityId: Number(formData.cityId),
      provinceId: Number(formData.provinceId),
      districtId: Number(formData.districtId),
      subdistrictId: Number(formData.subdistrictId),
      postalCodeId: Number(formData.postalCodeId),
    };
    await deleteAddress(address.id);
    invalidateUserAddresses();
  };

  return (
    <FormDelete
      loading={loading}
      onClick={handleDelete}
      title="Address Deletion"
      description="Are you sure want to delete this Address ?"
    />
  );
};

export default DeleteAddress;
