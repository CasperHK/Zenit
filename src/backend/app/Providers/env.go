package providers

import (
    "bufio"
    "os"
    "path/filepath"
    "strings"
)

// LoadEnv loads environment variables from a .env file in the working directory.
// It will not override variables that are already set in the environment.
func LoadEnv() error {
    // prefer .env in current working directory
    wd, err := os.Getwd()
    if err != nil {
        return err
    }

    files := []string{
        filepath.Join(wd, ".env"),
        filepath.Join(wd, ".env.local"),
    }

    for _, f := range files {
        if err := loadFileIfExists(f); err != nil {
            return err
        }
    }

    // If APP_ENV was set by the file or already present, attempt to load .env.<env>
    env := os.Getenv("APP_ENV")
    if env == "" {
        env = "development"
    }
    envFile := filepath.Join(wd, ".env."+env)
    _ = loadFileIfExists(envFile)
    return nil
}

func loadFileIfExists(path string) error {
    fi, err := os.Stat(path)
    if err != nil || fi.IsDir() {
        return nil
    }
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }
        // split at first '='
        parts := strings.SplitN(line, "=", 2)
        if len(parts) != 2 {
            continue
        }
        key := strings.TrimSpace(parts[0])
        val := strings.TrimSpace(parts[1])
        // remove surrounding quotes
        if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") && len(val) >= 2 {
            val = val[1 : len(val)-1]
        } else if strings.HasPrefix(val, "'") && strings.HasSuffix(val, "'") && len(val) >= 2 {
            val = val[1 : len(val)-1]
        }
        if os.Getenv(key) == "" {
            os.Setenv(key, val)
        }
    }
    return scanner.Err()
}
