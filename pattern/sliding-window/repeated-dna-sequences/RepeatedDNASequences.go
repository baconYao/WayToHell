package repeateddnasequences

func FindRepeatedSequences(s string, k int) Set {
	temp := *NewSet()
	output := *NewSet()
	for front := 0; front+k-1 < len(s); front++ {
		subString := s[front : front+k]
		if !temp.Exists(subString) {
			temp.Add(subString)
		} else {
			output.Add(subString)
		}
	}
	return output
}

/*
Adapted from: https://pkg.go.dev/github.com/deckarep/golang-set
Author's licensing terms:
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

type Set struct {
	hashMap map[interface{}]bool
}

// NewSet will initialize and return a new object of Set.
func NewSet() *Set {
	s := new(Set)
	s.hashMap = make(map[interface{}]bool)
	return s
}

// Add will add the value in the Set.
func (s *Set) Add(value interface{}) {
	s.hashMap[value] = true
}

// Delete will delete the value from the set.
func (s *Set) Delete(value interface{}) {
	delete(s.hashMap, value)
}

// Exists will check if the value exists in the set or not.
func (s *Set) Exists(value interface{}) bool {
	_, ok := s.hashMap[value]
	return ok
}
