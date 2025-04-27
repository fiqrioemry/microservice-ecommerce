// src/hooks/useProductMutation.js
import products from "@/services/products";
import { toast } from "sonner";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export const useProductMutation = () => {
  const queryClient = useQueryClient();

  const mutationOptions = (defaultMessage, refetchFn) => ({
    onSuccess: (response, variables) => {
      const successMessage = response?.message || defaultMessage;
      toast.success(successMessage);

      if (typeof refetchFn === "function") {
        refetchFn(variables);
      } else {
        queryClient.invalidateQueries({ queryKey: ["products"] });
      }
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
    createProduct: useMutation({
      mutationFn: products.onSuccesscreateProduct,
      ...mutationOptions("Product created successfully"),
    }),

    updateProduct: useMutation({
      mutationFn: ({ id, formData }) =>
        products.onSuccessupdateProduct(id, formData),
      ...mutationOptions("Product updated successfully", ({ id }) => {
        queryClient.invalidateQueries({ queryKey: ["product", id] });
        queryClient.invalidateQueries({ queryKey: ["products"] });
      }),
    }),

    deleteProduct: useMutation({
      mutationFn: (id) => products.onSuccessdeleteProduct(id),
      ...mutationOptions("Product deleted successfully", () => {
        queryClient.invalidateQueries({ queryKey: ["products"] });
      }),
    }),

    deleteVariantProduct: useMutation({
      mutationFn: (id) => products.onSuccessdeleteVariantProduct(id),
      ...mutationOptions("Product variant deleted successfully", () => {
        queryClient.invalidateQueries({ queryKey: ["products"] });
      }),
    }),
  };
};
