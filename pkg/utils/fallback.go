package utils

/*
# This clause handles errors returned by Ecto's insert/update/delete.
def call(conn, {:error, %Ecto.Changeset{} = changeset}) do
conn
|> put_status(:unprocessable_entity)
|> put_view(json: SpajzaWeb.ChangesetJSON)
|> render(:error, changeset: changeset)
end

# This clause is an example of how to handle resources that cannot be found.
def call(conn, {:error, :not_found}) do
conn
|> put_status(:not_found)
|> put_view(html: SpajzaWeb.ErrorHTML, json: SpajzaWeb.ErrorJSON)
|> render(:"404")
end

*/
