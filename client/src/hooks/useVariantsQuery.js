// src/hooks/useVariantsQuery.js
import { useQuery } from "@tanstack/react-query";
import { getVariants } from "@/services/variants";

export const useVariantsQuery = () => {
  return useQuery({
    queryKey: ["variants"],
    queryFn: getVariants,
    staleTime: 5 * 60 * 1000, // cache 5 menit
    refetchOnWindowFocus: false,
  });
};
