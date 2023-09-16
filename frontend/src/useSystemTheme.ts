import React from "react";

/**
 * System theme media query
 */
export const prefersDarkTheme =
  typeof window !== "undefined"
    ? matchMedia("(prefers-color-scheme: dark)")
    : null;

// Hook to return the user's preferred theme using media query
export default function useSystemTheme() {
  const [theme, setTheme] = React.useState<"dark" | "light">(
    prefersDarkTheme?.matches ? "dark" : "light"
  );

  React.useEffect(() => {
    const themeChangeHandler = (e: MediaQueryListEvent) => {
      setTheme(e.matches ? "dark" : "light");
    };
    prefersDarkTheme?.addEventListener("change", themeChangeHandler);
    return () => {
      prefersDarkTheme?.removeEventListener("change", themeChangeHandler);
    };
  }, []);

  return theme;
}
