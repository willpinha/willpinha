using Octokit;

namespace Willpinha.Console;

public static class Requests
{
    public static readonly string Username = "willpinha";

    public static async Task<SearchRepositoryResult> GetFamousRepositories()
    {
        var request = new SearchRepositoriesRequest()
        {
            User = Username,
            SortField = RepoSearchSort.Stars
        };

        return await Client.Instance.Search.SearchRepo(request);
    }

    public static async Task<SearchIssuesResult> GetLatestPullRequests()
    {
        var request = new SearchIssuesRequest
        {
            Type = IssueTypeQualifier.PullRequest,
            Involves = Username,
            Order = SortDirection.Descending,
            SortField = IssueSearchSort.Updated,
            PerPage = 10,
        };

        return await Client.Instance.Search.SearchIssues(request);
    }

    public static async Task<SearchIssuesResult> GetLatestIssues()
    {
        var request = new SearchIssuesRequest
        {
            Type = IssueTypeQualifier.Issue,
            Involves = Username,
            Order = SortDirection.Descending,
            SortField = IssueSearchSort.Updated,
            PerPage = 10,
        };

        return await Client.Instance.Search.SearchIssues(request);
    }
}