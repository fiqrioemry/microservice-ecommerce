// src/hooks/useCategoryMutation.js
import { toast } from "sonner";
import cat from "@/services/categories";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export const useCategoryMutation = () => {
  const queryClient = useQueryClient();

  const mutationOptions = (defaultMessage, refetchFn) => ({
    onSuccess: (response, variables) => {
      const successMessage = response?.message || defaultMessage;
      toast.success(successMessage);

      if (typeof refetchFn === "function") {
        refetchFn(variables);
      } else {
        queryClient.invalidateQueries(["categories"]);
      }
    },
    onError: (error) => {
      const errorMessage =
        error?.response?.data?.message || "Something went wrong";
      toast.error(errorMessage);
    },
  });

  return {
    createCategory: useMutation({
      mutationFn: (data) => cat.createCategory(data),
      ...mutationOptions("Category created successfully"),
    }),

    updateCategory: useMutation({
      mutationFn: ({ id, data }) => cat.updateCategory(id, data),
      ...mutationOptions("Category updated successfully", ({ id }) => {
        queryClient.invalidateQueries(["category", id]);
        queryClient.invalidateQueries(["categories"]);
      }),
    }),

    deleteCategory: useMutation({
      mutationFn: (id) => cat.deleteCategory(id),
      ...mutationOptions("Category deleted successfully", () => {
        queryClient.invalidateQueries(["categories"]);
      }),
    }),

    createSubcategory: useMutation({
      mutationFn: ({ categoryId, data }) =>
        cat.createSubcategory(categoryId, data),
      ...mutationOptions(
        "Subcategory created successfully",
        ({ categoryId }) => {
          queryClient.invalidateQueries(["category", categoryId]);
          queryClient.invalidateQueries(["categories"]);
        }
      ),
    }),

    updateSubcategory: useMutation({
      mutationFn: ({ subId, data }) => cat.updateSubcategory(subId, data),
      ...mutationOptions("Subcategory updated successfully", ({ subId }) => {
        queryClient.invalidateQueries(["subcategory", subId]);
        queryClient.invalidateQueries(["categories"]);
      }),
    }),

    deleteSubcategory: useMutation({
      mutationFn: (subId) => cat.deleteSubcategory(subId),
      ...mutationOptions("Subcategory deleted successfully", () => {
        queryClient.invalidateQueries(["categories"]);
      }),
    }),
  };
};
