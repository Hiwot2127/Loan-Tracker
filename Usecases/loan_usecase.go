package Usecases

import (
    "context"
    "Loan-Tracker/Domain"
    "Loan-Tracker/Repositories"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanUsecase struct {
    loanRepository *Repositories.LoanRepository
}

func NewLoanUsecase(loanRepo *Repositories.LoanRepository) *LoanUsecase {
    return &LoanUsecase{
        loanRepository: loanRepo,
    }
}

func (u *LoanUsecase) ApplyForLoan(ctx context.Context, loan *Domain.Loan) error {
    return u.loanRepository.CreateLoan(ctx, loan)
}

func (u *LoanUsecase) GetLoanStatus(ctx context.Context, id primitive.ObjectID) (*Domain.Loan, error) {
    return u.loanRepository.GetLoanByID(ctx, id)
}

func (u *LoanUsecase) GetAllLoans(ctx context.Context, filter map[string]interface{}) ([]Domain.Loan, error) {
    return u.loanRepository.GetAllLoans(ctx, filter)
}

func (u *LoanUsecase) UpdateLoanStatus(ctx context.Context, id primitive.ObjectID, status string) error {
    return u.loanRepository.UpdateLoanStatus(ctx, id, status)
}

func (u *LoanUsecase) DeleteLoan(ctx context.Context, id primitive.ObjectID) error {
    return u.loanRepository.DeleteLoan(ctx, id)
}