// src/components/input/SelectCategoryElement.jsx
import { useFormContext, Controller } from "react-hook-form";
import { useCategoriesQuery } from "@/hooks/useCategoriesQuery";

export const SelectCategoryElement = () => {
  const { control, watch, setValue } = useFormContext();
  const selectedCategoryId = watch("categoryId");

  const { data: rawCategories = { categories: [] } } = useCategoriesQuery();
  const categories = rawCategories.categories || [];

  const selectedCategory = categories.find(
    (cat) => cat.ID === selectedCategoryId
  );
  const subcategories = selectedCategory?.Subcategories || [];

  const handleCategoryChange = (e, field) => {
    const newCategoryId = e.target.value;
    field.onChange(newCategoryId);
    setValue("subcategoryId", ""); // Reset subcategory kalau ganti kategori
  };

  return (
    <div className="space-y-4">
      {/* Category */}
      <Controller
        control={control}
        name="categoryId"
        rules={{ required: true }}
        render={({ field, fieldState }) => (
          <div className="space-y-1">
            <label className="block text-sm font-medium">Category</label>
            <select
              {...field}
              onChange={(e) => handleCategoryChange(e, field)}
              className="w-full border p-2 rounded disabled:bg-gray-100"
            >
              <option value="">Select Category</option>
              {categories.map((cat) => (
                <option key={cat.ID} value={cat.ID}>
                  {cat.name}
                </option>
              ))}
            </select>
            {fieldState.error && (
              <p className="text-red-500 text-xs mt-1">
                {fieldState.error.message || "Please select a category"}
              </p>
            )}
          </div>
        )}
      />

      {/* Subcategory */}
      {selectedCategoryId && subcategories.length > 0 && (
        <Controller
          control={control}
          name="subcategoryId"
          rules={{ required: true }}
          render={({ field, fieldState }) => (
            <div className="space-y-1">
              <label className="block text-sm font-medium">Subcategory</label>
              <select
                {...field}
                className="w-full border p-2 rounded disabled:bg-gray-100"
              >
                <option value="">Select Subcategory</option>
                {subcategories.map((sub) => (
                  <option key={sub.ID} value={sub.ID}>
                    {sub.name}
                  </option>
                ))}
              </select>
              {fieldState.error && (
                <p className="text-red-500 text-xs mt-1">
                  {fieldState.error.message || "Please select a subcategory"}
                </p>
              )}
            </div>
          )}
        />
      )}
    </div>
  );
};
