// src/components/category/AddCategory.jsx
import React from "react";
import { PlusCircle } from "lucide-react";
import { categorySchema } from "@/lib/schema";
import { FormDialog } from "@/components/form/FormDialog";
import { InputElement } from "@/components/input/InputElement";
import { UploadElement } from "@/components/input/UploadElement";
import { categoryState, subCategoryState } from "@/lib/constant";
import { useCategoryMutation } from "@/hooks/useCategoryMutation";

const AddCategory = () => {
  const { mutateAsync: createCategory, isLoading } = useCategoryMutation();

  return (
    <FormDialog
      loading={isLoading}
      state={categoryState}
      action={createCategory}
      schema={categorySchema}
      title="Add New Category"
      buttonText={
        <button className="btn btn-primary gap-4">
          <PlusCircle size={18} />
          <span>Category</span>
        </button>
      }
    >
      <InputElement
        name="name"
        label="Category"
        placeholder="Masukkan Nama Category"
      />
      <UploadElement name="image" label="Category Image" maxImages={1} />
    </FormDialog>
  );
};

export default AddCategory;
