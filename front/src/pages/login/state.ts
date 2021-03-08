export interface AuthState {
  isAuthorized: boolean;
}

export const initialState: AuthState = { isAuthorized: false };
