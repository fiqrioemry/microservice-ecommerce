// src/hooks/useVariantMutation.js
import { useMutation, useQueryClient } from "@tanstack/react-query";
import {
  createVariant,
  updateVariant,
  deleteVariant,
  addVariantValue,
  updateVariantValue,
  deleteVariantValue,
} from "@/services/variants";

export const useCreateVariant = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: createVariant,
    onSuccess: () => {
      queryClient.invalidateQueries(["variants"]);
    },
  });
};

export const useUpdateVariant = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }) => updateVariant(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries(["variants"]);
    },
  });
};

export const useDeleteVariant = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id) => deleteVariant(id),
    onSuccess: () => {
      queryClient.invalidateQueries(["variants"]);
    },
  });
};

export const useAddVariantValue = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }) => addVariantValue(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries(["variants"]);
    },
  });
};

export const useUpdateVariantValue = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: ({ valueId, data }) => updateVariantValue(valueId, data),
    onSuccess: () => {
      queryClient.invalidateQueries(["variants"]);
    },
  });
};

export const useDeleteVariantValue = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (valueId) => deleteVariantValue(valueId),
    onSuccess: () => {
      queryClient.invalidateQueries(["variants"]);
    },
  });
};
