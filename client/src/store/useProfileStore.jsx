import { toast } from "sonner";
import { create } from "zustand";
import { useProfileManagement } from "@/hooks/useProfileManagement";
import users from "../services/users";
import locations from "../services/locations";

export const useProfileStore = create((set) => ({
  loading: false,
  // ✅ Update Profile
  updateProfile: async (data) => {
    set({ loading: true });
    try {
      users;
      await users.updateProfile(data);
      toast.success("Profile updated successfully");
      const { invalidateUserProfile } = useProfileManagement();
      invalidateUserProfile();
    } catch (error) {
      toast.error(error?.response?.data?.message);
    } finally {
      set({ loading: false });
    }
  },

  // ✅ Alamat: Tambah
  addAddress: async (data) => {
    set({ loading: true });
    try {
      await locations.addAddress(data);
      toast.success("Address added");
      const { invalidateUserAddresses } = useProfileManagement();
      invalidateUserAddresses();
    } catch (error) {
      toast.error(error?.response?.data?.message);
    } finally {
      set({ loading: false });
    }
  },

  // ✅ Alamat: Update
  updateAddress: async (id, data) => {
    set({ loading: true });
    try {
      await locations.updateAddress(id, data);
      toast.success("Address updated");
      const { invalidateUserAddresses } = useProfileManagement();
      invalidateUserAddresses();
    } catch (error) {
      toast.error(error?.response?.data?.message);
    } finally {
      set({ loading: false });
    }
  },

  // ✅ Alamat: Delete
  deleteAddress: async (id) => {
    set({ loading: true });
    try {
      await locations.deleteAddress(id);
      toast.success("Address deleted");
      const { invalidateUserAddresses } = useProfileManagement();
      invalidateUserAddresses();
    } catch (error) {
      toast.error(error?.response?.data?.message);
    } finally {
      set({ loading: false });
    }
  },

  // ✅ Alamat: Set Main
  setMainAddress: async (id) => {
    set({ loading: true });
    try {
      await locations.setMainAddress(id);
      toast.success("Main address set");
      const { invalidateUserAddresses } = useProfileManagement();
      invalidateUserAddresses();
    } catch (error) {
      toast.error(error?.response?.data?.message);
    } finally {
      set({ loading: false });
    }
  },
}));
