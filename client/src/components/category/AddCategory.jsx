// src/components/category/AddCategory.jsx
import React, { useState } from "react";
import { PlusCircle } from "lucide-react";
import { categorySchema } from "@/lib/schema";
import { FormDialog } from "@/components/form/FormDialog";
import { InputElement } from "@/components/input/InputElement";
import { UploadElement } from "@/components/input/UploadElement";
import { categoryState } from "@/lib/constant";
import { useCategoryMutation } from "@/hooks/useCategoryMutation";
import { toast } from "sonner";
import { useCategoryStore } from "../../store/useCategoryStore";

const AddCategory = () => {
  const { loading, AddNewCategory } = useCategoryStore();

  return (
    <FormDialog
      loading={loading}
      state={categoryState}
      schema={categorySchema}
      title="Add NewCategory"
      action={AddNewCategory}
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
      <UploadElement
        name="image"
        isSingle
        maxImages={1}
        label="Category Image"
      />
    </FormDialog>
  );
};

export default AddCategory;
