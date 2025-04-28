import React from "react";
import { MapPin } from "lucide-react";
import ErrorDialog from "@/components/ui/ErrorDialog";
import { useAddressesQuery } from "@/hooks/useUserQuery";
import FetchLoading from "@/components/ui/FetchLoading";
import { Card, CardContent } from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import AddAddress from "@/components/address/AddAddress";
import UpdateAddress from "@/components/address/UpdateAddress";
import DeleteAddress from "@/components/address/DeleteAddress";

const Address = () => {
  const { data: addresses, isError, isLoading } = useAddressesQuery();

  if (isLoading) return <FetchLoading />;

  if (isError) return <ErrorDialog />;

  return (
    <div className="space-y-6">
      <div className="flex justify-end">
        <AddAddress />
      </div>

      {addresses.length === 0 ? (
        <div className="flex flex-col items-center justify-center h-96 text-center text-muted-foreground space-y-4">
          <MapPin className="w-16 h-16 text-primary" />
          <p className="text-lg font-semibold">No Address Found</p>
          <p className="text-sm">
            Let's add your first address for easier checkout!
          </p>
          <AddAddress />
        </div>
      ) : (
        <ScrollArea className="h-[500px]">
          {addresses.map((add) => (
            <Card
              key={add.id}
              className="border border-border mb-4 hover:shadow-md transition"
            >
              <CardContent className="p-5 space-y-3 ">
                <div className="flex justify-between items-start">
                  <div className="space-y-1">
                    <h4 className="text-base font-semibold capitalize">
                      {add.name}
                    </h4>
                    <p className="text-sm text-muted-foreground">{add.phone}</p>
                  </div>
                  {add.isMain && (
                    <span className="text-xs font-medium text-white bg-primary rounded-full px-3 py-1">
                      Main
                    </span>
                  )}
                </div>

                <div className="text-sm space-y-1">
                  <p>{add.address}</p>
                  <p className="text-muted-foreground">
                    {add.subdistrict}, {add.district}, {add.city},{" "}
                    {add.province} - {add.postalCode}
                  </p>
                </div>

                <div className="flex gap-3 pt-2">
                  <UpdateAddress address={add} />
                  <DeleteAddress address={add} />
                </div>
              </CardContent>
            </Card>
          ))}
        </ScrollArea>
      )}
    </div>
  );
};

export default Address;
