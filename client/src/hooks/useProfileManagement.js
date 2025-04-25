import user from "@/services/users";
import { useQuery, useQueryClient } from "@tanstack/react-query";

// fetch data profile user
export const useProfileQuery = () => {
  return useQuery({
    queryKey: ["user-profile"],
    queryFn: user.getProfile,
  });
};

// fetch daftar alamat user
export const useAddressesQuery = () => {
  return useQuery({
    queryKey: ["user-addresses"],
    queryFn: user.getAddresses,
  });
};

export const useProfileManagement = () => {
  const queryClient = useQueryClient();

  return {
    invalidateUserProfile: () =>
      queryClient.invalidateQueries(["user-profile"]),
    invalidateUserAddresses: () =>
      queryClient.invalidateQueries(["user-addresses"]),
  };
};
