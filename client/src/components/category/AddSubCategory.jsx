// src/components/category/AddSubCategory.jsx
import React from "react";
import { PlusCircle } from "lucide-react";
import { subCategorySchema } from "@/lib/schema";
import { subCategoryState } from "@/lib/constant";
import { FormDialog } from "@/components/form/FormDialog";
import { InputElement } from "@/components/input/InputElement";
import { UploadElement } from "@/components/input/UploadElement";
import { useCategoryMutation } from "@/hooks/useCategoryMutation";

const AddSubCategory = () => {
  const { mutateAsync: createSubcategory, isLoading } = useCategoryMutation();

  return (
    <FormDialog
      loading={isLoading}
      state={subCategoryState}
      action={createSubcategory}
      schema={subCategorySchema}
      title="Add New Subcategory"
      buttonText={
        <button className="btn btn-primary gap-4">
          <PlusCircle size={18} />
          <span>Subategory</span>
        </button>
      }
    >
      <InputElement
        name="name"
        label="Subcategory"
        placeholder="Masukkan Nama Subcategory"
      />
      <UploadElement name="image" label="Subcategory Image" maxImages={1} />
    </FormDialog>
  );
};

export default AddSubCategory;
