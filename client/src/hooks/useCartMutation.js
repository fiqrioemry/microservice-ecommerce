import { toast } from "sonner";
import cart from "@/services/carts";
import { useMutation, useQueryClient } from "@tanstack/react-query";

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

  const { mutate: addToCart, ...addToCartRest } = useMutation({
    mutationFn: cart.addToCart,
    ...mutationOptions("Item added to cart"),
  });

  const { mutate: updateCartItem, ...updateCartItemRest } = useMutation({
    mutationFn: ({ itemId, data }) => cart.updateCartItem(itemId, data),
    ...mutationOptions("Cart item updated"),
  });

  const { mutate: removeCartItem, ...removeCartItemRest } = useMutation({
    mutationFn: cart.removeCartItem,
    ...mutationOptions("Cart item removed"),
  });

  const { mutate: clearCart, ...clearCartRest } = useMutation({
    mutationFn: cart.clearCart,
    ...mutationOptions("Cart cleared"),
  });

  return {
    addToCart,
    updateCartItem,
    removeCartItem,
    clearCart,
    ...addToCartRest,
    ...updateCartItemRest,
    ...removeCartItemRest,
    ...clearCartRest,
  };
};
