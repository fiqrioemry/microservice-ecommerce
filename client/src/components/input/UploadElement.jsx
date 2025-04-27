import { useState } from "react";
import { PlusCircle, X } from "lucide-react";
import { Controller, useFormContext } from "react-hook-form";

const UploadElement = ({
  name,
  label,
  maxImages = 5,
  rules = { required: true },
}) => {
  const { control } = useFormContext();
  const [previews, setPreviews] = useState([]);

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field, fieldState }) => {
        const handleAddImage = (e) => {
          const files = Array.from(e.target.files);
          const newPreviews = files.map((file) => ({
            id: URL.createObjectURL(file),
            file,
          }));

          const updated = [...(field.value || []), ...newPreviews];
          if (updated.length > maxImages) return;
          field.onChange(updated);
          setPreviews(updated);
        };

        const handleRemoveImage = (id) => {
          const updated = (field.value || []).filter((img) => img.id !== id);
          field.onChange(updated);
          setPreviews(updated);
        };

        return (
          <div className="space-y-2">
            {label && (
              <label className="block text-sm font-medium text-gray-700">
                {label}
              </label>
            )}

            <div className="flex flex-wrap gap-4">
              {(field.value || []).map((img) => (
                <div
                  key={img.id}
                  className="relative w-32 h-32 rounded-md overflow-hidden border border-border"
                >
                  <img
                    src={img.id}
                    alt="preview"
                    className="object-cover w-full h-full"
                  />
                  <button
                    type="button"
                    className="absolute top-1 right-1 bg-white rounded-full p-1 shadow-md hover:bg-red-500 hover:text-white transition"
                    onClick={() => handleRemoveImage(img.id)}
                  >
                    <X className="w-4 h-4" />
                  </button>
                </div>
              ))}

              {(!field.value || field.value.length < maxImages) && (
                <>
                  <label
                    htmlFor={`${name}-upload`}
                    className="flex items-center justify-center w-32 h-32 border-2 border-dashed border-primary rounded-md cursor-pointer hover:bg-primary/10 transition"
                  >
                    <span className="text-primary font-medium text-sm">
                      <PlusCircle />
                    </span>
                  </label>
                  <input
                    id={`${name}-upload`}
                    type="file"
                    accept="image/*"
                    onChange={handleAddImage}
                    className="hidden"
                  />
                </>
              )}
            </div>

            {fieldState.error && (
              <p className="text-red-500 text-xs mt-1">
                {fieldState.error.message || "This field is required"}
              </p>
            )}
          </div>
        );
      }}
    />
  );
};

export { UploadElement };
