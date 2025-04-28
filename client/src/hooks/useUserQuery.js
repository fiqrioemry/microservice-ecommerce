// src/hooks/useUserQueries.js

import userService from "@/services/users";
import { useQuery } from "@tanstack/react-query";

// ======== PROFILE QUERY ========
export const useProfileQuery = () =>
  useQuery({
    queryKey: ["profile"],
    queryFn: userService.getProfile,
  });

// ======== ADDRESS QUERY ========
export const useAddressesQuery = () =>
  useQuery({
    queryKey: ["addresses"],
    queryFn: userService.getAddresses,
  });
