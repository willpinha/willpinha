using Octokit;

namespace Willpinha.Console;

// HTTP client singleton
public static class Client
{
    public static GitHubClient Instance { get; } = new(new ProductHeaderValue("willpinha-profile-readme"))
    {
        Credentials = new Credentials(
            // Personal Access Token (PAT) from GitHub
            Environment.GetEnvironmentVariable("GITHUB_TOKEN") ?? throw new InvalidOperationException("GITHUB_TOKEN environment variable is not defined")
        )
    };
}