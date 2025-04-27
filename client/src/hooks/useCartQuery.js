// src/hooks/useCartQuery.js
import { useQuery } from "@tanstack/react-query";
import cart from "@/services/carts";

export const useCartQuery = () => {
  return useQuery({
    queryKey: ["cart"],
    queryFn: cart.getCart,
    staleTime: 1000 * 60 * 5,
    retry: 1,
  });
};
