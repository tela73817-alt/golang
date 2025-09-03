## Bài 7: Student Grade Manager

### Yêu cầu:
Tạo hệ thống quản lý điểm sinh viên với các tính năng:

**Struct Student:**
- Name (string)
- Grades ([]float64)

**Methods cần implement:**
- `NewStudent(name string) *Student`
- `AddGrade(grade float64) error` - Thêm điểm (0-100)
- `GetAverage() float64` - Tính điểm trung bình
- `GetLetterGrade() string` - Chuyển đổi sang điểm chữ (A, B, C, D, F)
- `GetHighestGrade() float64`
- `GetLowestGrade() float64`
- `GetGradeCount() int`
- `String() string`

**Business Rules:**
- Điểm phải từ 0-100
- Letter Grade: A(90-100), B(80-89), C(70-79), D(60-69), F(<60)
- Trả về 0 nếu chưa có điểm nào

**Test Cases:**
- Test thêm điểm hợp lệ/không hợp lệ
- Test tính trung bình với nhiều trường hợp
- Test letter grade conversion
- Test edge cases (no grades, single grade)

---