// src/hooks/useCartQuery.js
import cart from "@/services/carts";
import { useQuery } from "@tanstack/react-query";

export const useCartQuery = () => {
  return useQuery({
    queryKey: ["cart"],
    queryFn: cart.getCart,
    staleTime: 1000 * 60 * 5, // cache 5 menit
    retry: 1, // 1 x retry
  });
};
