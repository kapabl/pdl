namespace Io.Pdl.Infrastructure.Data;

public readonly record struct Filter(string Column, Operator Operator, object? Value);
