package branch

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

type GitRepository struct {
	// Repository holds the  git repository
	Repository *git.Repository
}

type References struct {
	// Refs holds the git repository's references
	Refs storer.ReferenceIter
}

type RefName struct {
  // S is the short hand name for the reference name
  S string
  
  // P is the full reference name. This is passed to a Repository.Storer's RemoveReference Method. 
  P plumbing.ReferenceName

  // Remote is a bool to tell if the reference is a remote or not
  Remote bool
}

// NewGitRepositoryFromString accepts a path to a git repository and
// returns a pointer to a RepositoryBranches object
func NewGitRepositoryFromString(s string) (*GitRepository, error) {

	repo, err := git.PlainOpen(s)
	if err != nil {
		return nil, err
	}

	return &GitRepository{
		Repository: repo,
	}, nil

}

// NewReferences accepts a GitRepository and returns a pointer to a
// References object
func NewReferences(g *GitRepository) (*References, error) {

	refs, err := g.Repository.References()
	if err != nil {
		return nil, err
	}
	return &References{
		Refs: refs,
	}, nil
}

// GetReferenceNames gets all of the short hand reference names from the
// current git repository
func (r *References) GetReferenceNames() ([]RefName, error) {
	var refNames []RefName

	r.Refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Type() == plumbing.HashReference && !ref.Name().IsTag() &&!ref.Name().IsRemote() {
                        currRef := &RefName{
                          S: ref.Name().Short(),
                          P: ref.Name(),
                          Remote: false,
                        }
			refNames = append(refNames, *currRef)
		}
		return nil
	})
	return refNames, nil
}

func (r *References) GetReferenceNamesWithRemotes() ([]RefName, error) {
	var refNames []RefName

	r.Refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Type() == plumbing.HashReference && !ref.Name().IsTag() {
                  if !ref.Name().IsRemote() {
                    currRef := &RefName{
                      S: ref.Name().Short(),
                      P: ref.Name(),
                      Remote: false,
                    }
		    refNames = append(refNames, *currRef)
                    } else {
                    currRef := &RefName{
                      S: ref.Name().Short(),
                      P: ref.Name(),
                      Remote: true,
                    }
                    refNames = append(refNames, *currRef)
                  }
		}
		return nil
	})
	return refNames, nil
}

// GetReferenceHashes gets all of the hashes from the current git repository
func (r *References) GetReferenceHashes() ([]plumbing.Hash, error) {
	var refHashes []plumbing.Hash

	r.Refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Type() == plumbing.HashReference && !ref.Name().IsTag() {
			refHashes = append(refHashes, ref.Hash())
		}
		return nil
	})
	return refHashes, nil
}

// GetReferenceMap returns a map of reference names as keys and the associated
// hash as the value
func (r *References) GetReferenceMap() (map[string]plumbing.Hash, error) {
	var refsHashMap map[string]plumbing.Hash

	r.Refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Type() == plumbing.HashReference && !ref.Name().IsTag() {
			refsHashMap[ref.Name().Short()] = ref.Hash()
		}
		return nil
	})

	return refsHashMap, nil
}
