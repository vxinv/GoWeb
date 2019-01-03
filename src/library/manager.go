package library
import "errors"
type Music struct {
	Id string
    Name string
    Artist string
    Source string
    Type string
}
type MusicManager struct {
	Musics []Music
}
func NewMusicManager() *MusicManager  {
	return &MusicManager{make([]Music,0)}
}
func (m *MusicManager)Len () int {
	return len(m.Musics)
}
func (m *MusicManager)Get(index int)(music *Music,err error){
	if(index<0||index>len(m.Musics)){
		return nil,errors.New("index out of range")
	}
	return &m.Musics[index],nil
}
func (m *MusicManager)Add(music *Music){
	m.Musics=append(m.Musics,*music)
}
func (m *MusicManager)Find(name string) *Music{
	if(len(m.Musics)==0){
		return nil
	}
	for _,m :=range m.Musics{
		if m.Name==name{
			return &m
		}
	}
	return nil
}
func (m *MusicManager)Remove(index int)*Music  {
   if(index<0||index>len(m.Musics)){
   	return nil
   }
   reMoveMusic:=&m.Musics[index];
   if(index<len(m.Musics)-1){
   	  m.Musics=append(m.Musics[:index-1],m.Musics[index+1:]...)
   	  //copy(m.Musics[index:],m.Musics[index+1:])
   }else if index==0{
   	 m.Musics=make([]Music ,0)
   }else {
   	m.Musics=m.Musics[:index-1]
   }
   return reMoveMusic
}

