package auth

import (
	"context"
	"net/http"

	pb "rj/tago/auth/proto"

	"google.golang.org/api/oauth2/v2"
)

var httpClient = &http.Client{}

// VerifyIDToken verify google auth
func verifyIDToken(idToken string) (*oauth2.Tokeninfo, error) {
	oauth2Service, err := oauth2.New(httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

// Server class
type Server struct {
	//...
}

// ThirdPartySignIn request response
// third party sign in, id token checking
func (s *Server) ThirdPartySignIn(ctx context.Context, request *pb.ThirdPartyAuth) (*pb.AuthResponse, error) {
	if request.Provider == pb.ThirdPartyAuth_GOOGLE {
		tokenInfo, err := verifyIDToken(request.IdToken)
		if err != nil {
			return nil, err
		}
		println(tokenInfo)
		return &pb.AuthResponse{Result: "Success "}, nil
	}
	return &pb.AuthResponse{Result: "Third Party Sign In Provider Not Implement Yet"}, nil
}

// TagoSignIn request response
// waiting implementation
func (s *Server) TagoSignIn(ctx context.Context, request *pb.TagoAuth) (*pb.AuthResponse, error) {
	return &pb.AuthResponse{Result: "Not implement yet"}, nil
}

// Register request response
// waiting implementation
func (s *Server) Register(ctx context.Context, request *pb.TagoRegister) (*pb.AuthResponse, error) {
	return &pb.AuthResponse{Result: "Not implement yet"}, nil
}

// NewServer return server
func NewServer() *Server {
	return &Server{}
}
