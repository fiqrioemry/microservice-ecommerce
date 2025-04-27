// src/hooks/useCategoryMutation.js
import cat from "@/services/categories";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export const useCategoryMutation = () => {
  const queryClient = useQueryClient();

  return {
    createCategory: useMutation({
      mutationFn: (data) => cat.createCategory(data),
      onSuccess: () => {
        queryClient.invalidateQueries(["categories"]);
      },
    }),

    updateCategory: useMutation({
      mutationFn: ({ id, data }) => cat.updateCategory(id, data),
      onSuccess: () => {
        queryClient.invalidateQueries(["categories"]);
      },
    }),

    deleteCategory: useMutation({
      mutationFn: (id) => cat.deleteCategory(id),
      onSuccess: () => {
        queryClient.invalidateQueries(["categories"]);
      },
    }),

    createSubcategory: useMutation({
      mutationFn: ({ categoryId, data }) =>
        cat.createSubcategory(categoryId, data),
      onSuccess: () => {
        queryClient.invalidateQueries(["categories"]);
      },
    }),

    updateSubcategory: useMutation({
      mutationFn: ({ subId, data }) => cat.updateSubcategory(subId, data),
      onSuccess: () => {
        queryClient.invalidateQueries(["categories"]);
      },
    }),

    deleteSubcategory: useMutation({
      mutationFn: (subId) => cat.deleteSubcategory(subId),
      onSuccess: () => {
        queryClient.invalidateQueries(["categories"]);
      },
    }),
  };
};
