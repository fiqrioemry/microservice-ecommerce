// src/hooks/useCartMutation.js
import { useMutation, useQueryClient } from "@tanstack/react-query";
import cart from "@/services/cart";
import { toast } from "sonner";

export const useCartMutation = () => {
  const queryClient = useQueryClient();

  const mutationOptions = (defaultMessage) => ({
    onSuccess: (response) => {
      toast.success(response?.message || defaultMessage);
      queryClient.invalidateQueries(["cart"]);
    },
    onError: (error) => {
      const errorMessage =
        error?.response?.data?.message ||
        error?.message ||
        "Something went wrong";
      toast.error(errorMessage);
    },
  });

  return {
    addToCart: useMutation({
      mutationFn: (data) => cart.addToCart(data),
      ...mutationOptions("Item added to cart"),
    }),

    updateCartItem: useMutation({
      mutationFn: ({ itemId, data }) => cart.updateCartItem(itemId, data),
      ...mutationOptions("Cart item updated"),
    }),

    removeCartItem: useMutation({
      mutationFn: (itemId) => cart.removeCartItem(itemId),
      ...mutationOptions("Cart item removed"),
    }),

    clearCart: useMutation({
      mutationFn: () => cart.clearCart(),
      ...mutationOptions("Cart cleared"),
    }),
  };
};
