import { useFormContext, Controller } from "react-hook-form";
import { useCategoriesQuery } from "@/hooks/useCategoriesQuery";

const CategorySelection = () => {
  const { control, watch } = useFormContext();
  const selectedCategoryId = watch("categoryId");

  const { data = {} } = useCategoriesQuery();
  const categories = data.categories || [];

  const selectedCategory = categories.find(
    (cat) => cat.ID === selectedCategoryId
  );

  const subcategories = selectedCategory?.Subcategories || [];

  return (
    <div className="space-y-4">
      {/* Category */}
      <Controller
        control={control}
        name="categoryId"
        rules={{ required: true }}
        render={({ field }) => (
          <div>
            <label className="block mb-1 font-medium">Category</label>
            <select {...field} className="w-full border p-2 rounded">
              <option value="">Select Category</option>
              {categories.map((cat) => (
                <option key={cat.ID} value={cat.ID}>
                  {cat.name}
                </option>
              ))}
            </select>
          </div>
        )}
      />

      {/* Subcategory */}
      {selectedCategoryId && (
        <Controller
          control={control}
          name="subcategoryId"
          rules={{ required: true }}
          render={({ field }) => (
            <div>
              <label className="block mb-1 font-medium">Subcategory</label>
              <select {...field} className="w-full border p-2 rounded">
                <option value="">Select Subcategory</option>
                {subcategories.map((sub) => (
                  <option key={sub.ID} value={sub.ID}>
                    {sub.name}
                  </option>
                ))}
              </select>
            </div>
          )}
        />
      )}
    </div>
  );
};

export default CategorySelection;
