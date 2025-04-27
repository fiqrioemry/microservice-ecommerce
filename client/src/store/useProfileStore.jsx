// src/store/useProfileStore.jsx
import { toast } from "sonner";
import { create } from "zustand";
import users from "@/services/users";

export const useProfileStore = create((set) => {
  const setLoading = (value) => set({ loading: value });

  const handleMutation = async (fn, successMessage) => {
    setLoading(true);
    try {
      const response = await fn();
      toast.success(response?.message || successMessage);
    } catch (error) {
      const errorMessage =
        error?.response?.data?.message ||
        error?.message ||
        "Something went wrong";
      toast.error(errorMessage);
    } finally {
      setLoading(false);
    }
  };

  return {
    loading: false,

    addAddress: (data) =>
      handleMutation(
        () => users.addAddress(data),
        "Address added successfully"
      ),

    updateAddress: (id, data) =>
      handleMutation(
        () => users.updateAddress(id, data),
        "Address updated successfully"
      ),

    deleteAddress: (id) =>
      handleMutation(
        () => users.deleteAddress(id),
        "Address deleted successfully"
      ),

    setMainAddress: (id) =>
      handleMutation(
        () => users.setMainAddress(id),
        "Main address set successfully"
      ),
  };
});
