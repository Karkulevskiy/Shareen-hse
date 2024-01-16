using MediatR;

namespace Shareen.Application.Lobbies.Queries.GetLobbiesList;

public class GetLobbyListQuery : IRequest<LobbiesListVm> { }