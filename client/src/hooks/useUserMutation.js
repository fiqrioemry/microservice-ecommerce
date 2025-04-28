// src/hooks/useUserMutations.js

import userService from "@/services/users";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";

// ======== PROFILE MUTATION ========
export const useUpdateProfileMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: userService.updateProfile,
    onSuccess: ({ message }) => {
      toast.success(message);
      queryClient.invalidateQueries({ queryKey: ["profile"] });
    },
    onError: (error) => {
      toast.error(error?.response?.data?.message || "Failed to update profile");
    },
  });
};

// ======== ADDRESS MUTATIONS ========
export const useAddAddressMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: userService.addAddress,
    onSuccess: ({ message }) => {
      toast.success(message);
      queryClient.invalidateQueries({ queryKey: ["addresses"] });
    },
    onError: (error) => {
      toast.error(error?.response?.data?.message || "Failed to add address");
    },
  });
};

// src/hooks/useUserMutation.js
export const useUpdateAddressMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }) => userService.updateAddress(id, data),
    onSuccess: ({ message }) => {
      toast.success(message);
      queryClient.invalidateQueries({ queryKey: ["addresses"] });
    },
    onError: (error) => {
      toast.error(error?.response?.data?.message || "Failed to update address");
    },
  });
};

export const useDeleteAddressMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id) => userService.deleteAddress(id),
    onSuccess: ({ message }) => {
      toast.success(message);
      queryClient.invalidateQueries({ queryKey: ["addresses"] });
    },
    onError: (error) => {
      toast.error(error?.response?.data?.message || "Failed to delete address");
    },
  });
};

export const useSetMainAddressMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id) => userService.setMainAddress(id),
    onSuccess: ({ message }) => {
      toast.success(message);
      queryClient.invalidateQueries({ queryKey: ["addresses"] });
    },
    onError: (error) => {
      toast.error(
        error?.response?.data?.message || "Failed to set main address"
      );
    },
  });
};
