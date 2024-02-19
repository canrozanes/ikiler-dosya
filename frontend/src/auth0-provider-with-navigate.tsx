import { Auth0Provider } from "@auth0/auth0-react";
import { useNavigate } from "react-router-dom";
import { getAuthOConfig } from "./auth0-config";

interface AuthProviderWithNavigateProps {
  children: React.ReactNode;
}

export const Auth0ProviderWithNavigate = ({
  children,
}: AuthProviderWithNavigateProps) => {
  const navigate = useNavigate();

  const authConfig = getAuthOConfig();

  const onRedirectCallback = (appState: any) => {
    navigate(appState?.returnTo || window.location.pathname);
  };

  if (
    !(
      authConfig.domain &&
      authConfig.clientId &&
      authConfig.redirectUri &&
      authConfig.audience
    )
  ) {
    return null;
  }

  return (
    <Auth0Provider
      domain={authConfig.domain}
      clientId={authConfig.clientId}
      authorizationParams={{
        audience: authConfig.audience,
        redirect_uri: authConfig.redirectUri,
      }}
      onRedirectCallback={onRedirectCallback}
    >
      {children}
    </Auth0Provider>
  );
};
